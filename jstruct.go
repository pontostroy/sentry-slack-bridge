package main

type Config struct {
	Name    string `yaml:"name"`
	Options struct {
		Path    string `yaml:"path"`
		Channel string `yaml:"channel"`
		Webhook string `yaml:"webhook"`
	} `yaml:"options"`
}

type WebHook struct {
	Culprit string `json:"culprit"`
	Event   struct {
		Metrics struct {
			Bytes_ingested_event int64 `json:"bytes.ingested.event"`
			Bytes_stored_event   int64 `json:"bytes.stored.event"`
		} `json:"_metrics"`
		Contexts struct {
			Runtime struct {
				Build   string `json:"build"`
				Name    string `json:"name"`
				Type    string `json:"type"`
				Version string `json:"version"`
			} `json:"runtime,omitempty"`
		} `json:"contexts,omitempty"`
		Culprit     string `json:"culprit"`
		Environment string `json:"environment"`
		EventID     string `json:"event_id"`
		Exception   struct {
			Values []struct {
				Mechanism struct {
					Data        interface{} `json:"data"`
					Description interface{} `json:"description"`
					Handled     bool        `json:"handled"`
					HelpLink    interface{} `json:"help_link"`
					Meta        interface{} `json:"meta"`
					Synthetic   interface{} `json:"synthetic"`
					Type        string      `json:"type"`
				} `json:"mechanism,omitempty"`
				Stacktrace struct {
					Frames []struct {
						AbsPath     string      `json:"abs_path"`
						Colno       interface{} `json:"colno"`
						ContextLine string      `json:"context_line"`
						Data        struct {
							OrigInApp int64 `json:"orig_in_app"`
						} `json:"data"`
						Errors          interface{} `json:"errors"`
						Filename        string      `json:"filename"`
						Function        string      `json:"function"`
						ImageAddr       interface{} `json:"image_addr"`
						InApp           bool        `json:"in_app"`
						InstructionAddr interface{} `json:"instruction_addr"`
						Lineno          int64       `json:"lineno"`
						Module          string      `json:"module"`
						Package         interface{} `json:"package"`
						Platform        interface{} `json:"platform"`
						PostContext     []string    `json:"post_context"`
						PreContext      []string    `json:"pre_context"`
						RawFunction     interface{} `json:"raw_function"`
						Symbol          interface{} `json:"symbol"`
						SymbolAddr      interface{} `json:"symbol_addr"`
						Trust           interface{} `json:"trust"`
						Vars            struct {
							Context string `json:"context"`
							Event   struct {
								Key1 string `json:"key1"`
								Key2 string `json:"key2"`
								Key3 string `json:"key3"`
							} `json:"event"`
							UserAttrs struct{} `json:"user_attrs"`
						} `json:"vars"`
					} `json:"frames"`
				} `json:"stacktrace"`
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"values"`
		} `json:"exception"`
		Extra struct {
			Cloudwatch_logs struct {
				LogGroup  string `json:"log_group"`
				LogStream string `json:"log_stream"`
				URL       string `json:"url"`
			} `json:"cloudwatch logs,omitempty"`
			Lambda struct {
				AwsRequestID          string `json:"aws_request_id"`
				FunctionName          string `json:"function_name"`
				FunctionVersion       string `json:"function_version"`
				InvokedFunctionArn    string `json:"invoked_function_arn"`
				RemainingTimeInMillis int64  `json:"remaining_time_in_millis"`
			} `json:"lambda,omitempty"`
			Sys_argv []string `json:"sys.argv"`
		} `json:"extra,omitempty"`
		Fingerprint    []string `json:"fingerprint"`
		GroupingConfig struct {
			Enhancements string `json:"enhancements"`
			ID           string `json:"id"`
		} `json:"grouping_config,omitempty"`
		Hashes   []string `json:"hashes"`
		ID       string   `json:"id"`
		KeyID    string   `json:"key_id"`
		Level    string   `json:"level"`
		Location string   `json:"location"`
		Logger   string   `json:"logger"`
		Metadata struct {
			Filename string `json:"filename"`
			Function string `json:"function"`
			Type     string `json:"type"`
			Value    string `json:"value"`
		} `json:"metadata,omitempty"`
		Modules struct {
			Boto3              string `json:"boto3"`
			Botocore           string `json:"botocore"`
			Certifi            string `json:"certifi"`
			Chardet            string `json:"chardet"`
			Docutils           string `json:"docutils"`
			Idna               string `json:"idna"`
			Jmespath           string `json:"jmespath"`
			Pip                string `json:"pip"`
			Python_dateutil    string `json:"python-dateutil"`
			Python_json_logger string `json:"python-json-logger"`
			S3transfer         string `json:"s3transfer"`
			Schematics         string `json:"schematics"`
			Sentry_sdk         string `json:"sentry-sdk"`
			Setuptools         string `json:"setuptools"`
			Six                string `json:"six"`
			Urllib3            string `json:"urllib3"`
		} `json:"modules,omitempty"`
		Platform string  `json:"platform"`
		Project  int64   `json:"project"`
		Received float64 `json:"received"`
		Request  struct {
			Cookies             interface{}   `json:"cookies"`
			Data                interface{}   `json:"data"`
			Env                 interface{}   `json:"env"`
			Fragment            interface{}   `json:"fragment"`
			Headers             interface{}   `json:"headers"`
			InferredContentType interface{}   `json:"inferred_content_type"`
			Method              interface{}   `json:"method"`
			QueryString         []interface{} `json:"query_string"`
			URL                 string        `json:"url"`
		} `json:"request,omitempty"`
		Sdk struct {
			Integrations []string `json:"integrations"`
			Name         string   `json:"name"`
			Packages     []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"packages"`
			Version string `json:"version"`
		} `json:"sdk,omitempty"`
		Tags        [][]string `json:"tags"`
		Timestamp   float64    `json:"timestamp"`
		Title       string     `json:"title"`
		Transaction string     `json:"transaction"`
		Type        string     `json:"type"`
		Version     string     `json:"version"`
	} `json:"event,omitempty"`
	ID              string      `json:"id,omitempty"`
	Level           string      `json:"level,omitempty"`
	Logger          interface{} `json:"logger,omitempty"`
	Message         string      `json:"message,omitempty"`
	Project         string      `json:"project,omitempty"`
	ProjectName     string      `json:"project_name,omitempty"`
	ProjectSlug     string      `json:"project_slug,omitempty"`
	TriggeringRules []string    `json:"triggering_rules,omitempty"`
	URL             string      `json:"url,omitempty"`
}
