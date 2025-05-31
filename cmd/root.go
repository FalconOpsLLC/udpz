package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/FalconOpsLLC/udpz/internal/data"
	"github.com/FalconOpsLLC/udpz/internal/scan"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	// Scan options
	hostConcurrency uint = 255
	portConcurrency uint = 100
	timeoutMs       uint = 3000
	retransmissions uint = 2

	// DNS options
	scanAllAddresses bool = true

	// Logging options
	quiet  bool = false // Disable info logging output (non-errors)
	silent bool = false // Disable logging entirely

	info  bool = true // Default log level
	debug bool = false
	trace bool = false

	// Probe selection options
	listServices bool
	useProbes    string
	useTags      string

	// Output options
	outputPath   string
	logPath      string
	outputFormat string = "auto"
	logFormat    string = "auto"
	outputAppend bool   = true

	// Proxy options
	socks5Address  string
	socks5User     string
	socks5Password string
	socks5Timeout  uint = 3000

	// Constraints
	supportedLogFormats = map[string]bool{
		"json": true, "jsonl": true,
		"pretty": true,
		"auto":   true,
	}
	supportedOutputFormats = map[string]bool{
		"text": true, "txt": true,
		"yaml": true, "yml": true,
		"json": true, "jsonl": true,
		"csv":    true,
		"tsv":    true,
		"pretty": true,
		"auto":   true,
	}
)

func init() {

	rootCmd.Flags().SortFlags = false
	rootCmd.InitDefaultVersionFlag()
	rootCmd.InitDefaultCompletionCmd()

	// Output
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", outputPath, "Save output to file")
	rootCmd.Flags().StringVarP(&logPath, "log", "O", logPath, "Output log messages to file")
	rootCmd.Flags().BoolVar(&outputAppend, "append", outputAppend, "Append results to output file")
	rootCmd.Flags().StringVarP(&outputFormat, "format", "f", outputFormat, "Output format [text, pretty, csv, tsv, json, yaml, auto]")
	rootCmd.Flags().StringVarP(&logFormat, "log-format", "F", logFormat, `Output log format [pretty, json, auto]`)

	// Probe Selection
	rootCmd.Flags().BoolVarP(&listServices, "list", "l", listServices, "List available services / probes")
	rootCmd.Flags().StringVarP(&useProbes, "probes", "p", useProbes, "comma-delimited list of service probes")
	rootCmd.Flags().StringVar(&useTags, "tags", useTags, "comma-delimited list of target service tags")

	// Performance
	rootCmd.Flags().UintVarP(&hostConcurrency, "host-tasks", "c", hostConcurrency, "Maximum Number of hosts to scan concurrently")
	rootCmd.Flags().UintVarP(&portConcurrency, "port-tasks", "P", portConcurrency, "Number of Concurrent scan tasks per host")
	rootCmd.Flags().UintVarP(&retransmissions, "retries", "r", retransmissions, "Number of probe retransmissions per probe")
	rootCmd.Flags().UintVarP(&timeoutMs, "timeout", "t", timeoutMs, "UDP Probe timeout in milliseconds")

	// DNS
	rootCmd.Flags().BoolVarP(&scanAllAddresses, "all", "A", scanAllAddresses, "Scan all resolved addresses instead of just the first")

	// Proxy
	rootCmd.Flags().StringVarP(&socks5Address, "socks", "S", socks5Address, "SOCKS5 proxy address as HOST:PORT")
	rootCmd.Flags().StringVar(&socks5User, "socks-user", socks5User, "SOCKS5 proxy username")
	rootCmd.Flags().StringVar(&socks5Password, "socks-pass", socks5Password, "SOCKS5 proxy password (WARNING: insecure)")
	rootCmd.Flags().UintVar(&socks5Timeout, "socks-timeout", socks5Timeout, "SOCKS5 proxy timeout in milliseconds")

	// Logging
	rootCmd.Flags().BoolVarP(&debug, "debug", "D", debug, "Enable debug logging (Noisy!)")
	rootCmd.Flags().BoolVarP(&trace, "trace", "T", trace, "Enable trace logging (Very noisy!)")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", quiet, "Disable info logging")
	rootCmd.Flags().BoolVarP(&silent, "silent", "s", silent, "Disable ALL logging")
}

var rootCmd = &cobra.Command{
	Use:     "udpz [flags] [IP|hostname|CIDR|file ...]",
	Short:   "Speedy probe-oriented UDP port scanner",
	Version: "0.1.3",
	Long: `
  ┳┳  ┳┓  ┏┓  ┏┓
  ┃┃━━┃┃━━┃┃━━┏┛
  ┗┛  ┻┛  ┣┛  ┗┛

  Author: Bryan McNulty (@bryanmcnulty)
  Source: https://github.com/FalconOpsLLC/udpz`,

	RunE: func(cmd *cobra.Command, targets []string) (err error) {

		var outputFile *os.File
		var log zerolog.Logger
		var logFile *os.File
		var outputFlags int = os.O_WRONLY | os.O_CREATE
		var logFileFlags int = os.O_WRONLY | os.O_CREATE | os.O_APPEND

		outputFormat = strings.ToLower(outputFormat)

		if len(targets) == 0 && !listServices {
			return errors.New("at least one target (IP address, hostname, CIDR, or file) is required")
		}
		if sup, ok := supportedOutputFormats[outputFormat]; !ok || !sup {
			return errors.New("invalid output format: " + outputFormat)
		}
		if sup, ok := supportedLogFormats[logFormat]; !ok || !sup {
			return errors.New("invalid log format: " + logFormat)
		}
		if portConcurrency < 1 || hostConcurrency < 1 {
			return errors.New("concurrency value must be > 0")
		}
		if timeoutMs < 1 {
			return errors.New("timeout value must be > 0")
		}
		if outputAppend {
			outputFlags |= os.O_APPEND
		}

		if silent {
			zerolog.SetGlobalLevel(zerolog.Disabled)
		} else if quiet {
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		} else if trace {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		} else if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else if info {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano

		if logPath == "" {
			log = zerolog.New(os.Stderr)

			if logFormat == "auto" || logFormat == "pretty" {
				log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			}

		} else if logFile, err = os.OpenFile(logPath, logFileFlags, 0o644); err == nil {
			defer logFile.Close()

			log = zerolog.New(logFile)

			if logFormat == "pretty" {
				log = log.Output(zerolog.ConsoleWriter{Out: logFile})
			}
		} else {
			log = zerolog.New(os.Stderr)

			if logFormat == "auto" || logFormat == "pretty" {
				log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			}
			log.Error().
				Stack().Err(err).
				Str("log_path", logPath).
				Msg("Could not open log file for writing")
		}
		log = log.
			With().
			Timestamp().
			Logger()

		broker := data.NewUdpDataBroker(log)

		if listServices {
			services, probes := broker.Filter(strings.Split(useProbes, ","), strings.Split(useTags, ","))
			log.Debug().
				Int("service_count", len(services)).
				Int("probe_count", len(probes)).
				Msg("Filtered services")

			if outputPath == "" {
				outputFile = os.Stdout

			} else if outputFile, err = os.OpenFile(outputPath, outputFlags, 0o644); err != nil {
				log.Error().
					AnErr("error", err).
					Str("outputPath", outputPath).
					Msg("Could not open output file for writing")
				outputFile = os.Stdout
			}
			if data, err := json.Marshal(services); err == nil {
				outputFile.Write(data)

			} else {
				log.Fatal().
					Str("path", outputPath).
					Err(err).
					Msg("Failed to serialize object")
			}
			return nil
		}

		var scanner scan.UdpProbeScanner

		if scanner, err = scan.NewUdpProbeScanner(
			&log,
			broker,
			scanAllAddresses,
			hostConcurrency,
			portConcurrency,
			retransmissions,
			time.Duration(timeoutMs)*time.Millisecond,
			socks5Address,
			socks5User,
			socks5Password,
			int(socks5Timeout*uint(time.Millisecond))); err != nil {

			log.Fatal().
				Err(err).
				Msg("Failed to initialize scanner")
		}

		var scanStartTime, scanEndTime time.Time

		scanStartTime = time.Now()
		if err = scanner.Scan(targets, strings.Split(useProbes, ","), strings.Split(useTags, ",")); err != nil {
			log.Fatal().
				Err(err).
				Msg("Scan aborted")
		}
		scanEndTime = time.Now()

		log.Info().
			Time("start", scanStartTime).
			Time("end", scanEndTime).
			TimeDiff("duration", scanEndTime, scanStartTime).
			Msg("Scan complete")

		if scanner.Length() > 0 {

			if outputPath == "" {
				outputFile = os.Stdout
				if outputFormat == "auto" {
					outputFormat = "pretty"
				}

			} else if outputFile, err = os.OpenFile(outputPath, outputFlags, 0o644); err == nil {
				if outputFormat == "auto" {
					outputFormat = "json"
				}
				defer outputFile.Close()

			} else {
				log.Error().
					AnErr("error", err).
					Str("outputPath", outputPath).
					Msg("Could not open output file for writing")
				outputFile = os.Stdout
				if outputFormat == "auto" {
					outputFormat = "pretty"
				}
			}
			if outputFormat == "json" || outputFormat == "jsonl" {
				scanner.SaveJson(outputFile)
			} else if outputFormat == "yml" || outputFormat == "yaml" {
				scanner.SaveYAML(outputFile)
			} else {
				scanner.SaveTable(outputFormat, outputFile)
			}
		}
		return
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
