package judge0

import(
	"time"
	"errors"
	"encoding/json"
	"encoding/base64"
	"strings"
)


type RequestFields struct{
	
	SourceCode string `json:"source_code"`
	LanguageID int `json:"language_id"`

	CompilerOptions *string `json:"compiler_options,omitempty"`
	CommandLineArguments *string `json:"command_line_arguments,omitempty"`
	Stdin *string `json:"stdin,omitempty"`
	ExpectedOutput *string `json:"expected_output,omitempty"`
	CPUTimeLimit *float64 `json:"cpu_time_limit,omitempty"`
	CPUExtraTime *float64 `json:"cpu_extra_time,omitempty"`
	WallTimeLimit *float64 `json:"wall_time_limit,omitempty"`
	MemoryLimit *float64 `json:"memory_limit,omitempty"`
	StackLimit	*int `json:"stack_limit,omitempty"`
	MaxProcessesAndOrThreads *int `json:"max_processes_and_or_threads,omitempty"`
	EnablePerProcessAndThreadTimeLimit *bool `json:"enable_per_process_and_thread_time_limit,omitempty"`
	EnablePerProcessAndThreadMemoryLimit *bool `json:"enable_per_process_and_thread_memory_limit,omitempty"`
	MaxFileSize *int `json:"max_file_size,omitempty"`
	RedirectStderrToStdout *bool `json:"redirect_stderr_to_stdout,omitempty"`
	EnableNetwork *bool `json:"enable_network,omitempty"`
	NumberOfRuns *int `json:"number_of_runs,omitempty"`
	AdditionalFiles *string `json:"additional_files,omitempty"`
	CallbackURL *string `json:"callback_url,omitempty"`

}

type ResponseFields struct{
	
	Stdout *string `json:"stdout"`
	Stderr *string `json:"stderr"`
	CompileOutput *string`json:"compile_output"`
	Message *string `json:"message"`
	ExitCode *int `json:"exit_code"`
	ExitSignal *int `json:"exit_signal"`
	Status *Status `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	FinishedAt *time.Time `json:"finished_at"`
	Token *string `json:"token"`
	Time *float64 `json:"time"`
	WallTime *float64 `json:"wall_time"`
	Memory *float64 `json:"memory"`

}


type Submission struct {

	Request RequestFields
	response ResponseFields
}




func encodeField(field *string){

	if field != nil{
		*field = base64.StdEncoding.EncodeToString([]byte(*field))
	}
}


func decodeField(field *string) error{
	if field != nil{
		temp, err := base64.StdEncoding.DecodeString(*field)
		if err!=nil{
			return err
		}
		*field=string(temp)
	}
	return nil
}






func (req *RequestFields) isValid() error{

	if strings.TrimSpace(req.SourceCode)==""{
		return errors.New("source_code is required")
	}

	if req.LanguageID<=0{ 
		return errors.New("invalid language_id")
	}

	return nil

}


func (request *RequestFields) asBody() ([]byte, error){

	req := *request

	err := req.isValid()
	if err!=nil{
		return nil, err
	}

	encodeField(&req.SourceCode)
	encodeField(req.Stdin)
	encodeField(req.AdditionalFiles)
	encodeField(req.ExpectedOutput)
	
	dat, err := json.Marshal(req)
	if err!=nil{
		return nil, err
	}


	return dat, nil
}






func (resp *ResponseFields) decodeResponse() (*ResponseFields, error){

	res := *resp

	var err error

	err = decodeField(res.Stdout)
	if err!=nil{
		return nil, err
	}
	
	err = decodeField(res.Stderr)
	if err!=nil{
		return nil, err
	}
	
	err = decodeField(res.CompileOutput)	
	if err!=nil{
		return nil, err
	}

	return &res, nil

}



