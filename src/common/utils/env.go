package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/goharbor/harbor/src/common/utils/log"
)

func InitEnvironment() error {
	if envFiles, ok := os.LookupEnv("ENV_FILES"); ok {
		for _, envFile := range strings.Split(envFiles, ",") {
			log.Infof("loading environment from %s\n", envFile)
			dotEnv, err := os.Open(strings.TrimSpace(envFile))
			if err != nil {
				return err
			}

			scanner := bufio.NewScanner(dotEnv)
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					continue
				}
				parts := strings.Split(line, "=")
				if len(parts) != 2 {
					return fmt.Errorf("Missing variable or value for %s", parts)
				}
				val := strings.Trim(parts[1], `"'`)

				err := os.Setenv(parts[0], val)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
