package cockpit

import (
	"regexp"

	"github.com/JulzDiverse/aviator"
	"github.com/JulzDiverse/aviator/evaluator"
	"github.com/JulzDiverse/aviator/executor"
	"github.com/JulzDiverse/aviator/filemanager"
	"github.com/JulzDiverse/aviator/processor"
	"github.com/JulzDiverse/aviator/squasher"
	"github.com/JulzDiverse/aviator/validator"
	"github.com/JulzDiverse/osenv"
	"github.com/pkg/errors"
	"github.com/starkandwayne/goutils/ansi"

	yaml "gopkg.in/yaml.v2"
)

type Cockpit struct {
	spruceProcessor aviator.SpruceProcessor
	flyExecutor     aviator.FlyExecuter
	validator       aviator.Validator
}

type Aviator struct {
	cockpit     *Cockpit
	AviatorYaml *aviator.AviatorYaml
}

func Init(
	spruceProcessor aviator.SpruceProcessor,
	flyExecuter aviator.FlyExecuter,
	validator aviator.Validator,
) *Cockpit {
	return &Cockpit{spruceProcessor, flyExecuter, validator}
}

func New(curlyBraces bool) *Cockpit {
	return &Cockpit{
		spruceProcessor: processor.New(curlyBraces),
		validator:       validator.New(),
		flyExecutor:     executor.NewFlyExecutor(),
	}
}

func (c *Cockpit) NewAviator(aviatorYml []byte, varsMap map[string]string) (*Aviator, error) {
	var aviator aviator.AviatorYaml
	aviatorYml, err := resolveEnvVars(aviatorYml)
	if err != nil {
		return nil, errors.Wrap(err, ansi.Sprintf("@R{Reading Failed}"))
	}

	aviatorYml, err = evaluator.Evaluate(aviatorYml, varsMap)
	if err != nil {
		return nil, err
	}

	aviatorYml = quoteCurlyBraces(aviatorYml)
	err = yaml.Unmarshal(aviatorYml, &aviator)
	if err != nil {
		return nil, errors.Wrap(err, ansi.Sprintf("@R{YAML Parsing Failed}"))
	}

	err = c.validator.ValidateSpruce(aviator.Spruce)
	if err != nil {
		return nil, err
	}

	return &Aviator{c, &aviator}, nil
}

func (a *Aviator) ProcessSprucePlan(verbose bool, silent bool) error {
	err := a.cockpit.spruceProcessor.ProcessWithOpts(a.AviatorYaml.Spruce, verbose, silent)
	if err != nil {
		return errors.Wrap(err, "Processing Spruce Plan FAILED")
	}
	return nil
}

func (a *Aviator) ProcessSquashPlan() error {
	var err error
	var result []byte

	store := filemanager.Store(false)
	fp := processor.FileProcessor{store}

	content := a.AviatorYaml.Squash.Content
	for _, c := range content {
		var squashed []byte
		if len(c.Files) != 0 {
			files := store.ReadFiles(c.Files)
			squashed, err = squasher.Squash(files)
		} else {
			paths := fp.CollectFilesFromDir(c.Dir, "", []string{})
			files := store.ReadFiles(paths)
			squashed, err = squasher.Squash(files)
		}

		if err != nil {
			return err
		}

		result = append(result, squashed...)
	}

	err = store.WriteFile(a.AviatorYaml.Squash.To, result)
	if err != nil {
		return err
	}

	return nil
}

func (a *Aviator) ExecuteFly() error {
	err := a.cockpit.flyExecutor.Execute(a.AviatorYaml.Fly)
	if err != nil {
		return err
	}
	return nil
}

func resolveEnvVars(input []byte) ([]byte, error) {
	result, err := osenv.ExpandEnv(string(input))
	return []byte(result), err
}

func quoteCurlyBraces(input []byte) []byte {
	quoteRegex := `(\{\{|\+\+)([-\_\.\/\w\p{L}\/]+)(\}\}|\+\+)`
	re := regexp.MustCompile("(" + quoteRegex + ")")
	return re.ReplaceAll(input, []byte("\"$1\""))
}
