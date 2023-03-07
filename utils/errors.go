package utils

import "errors"

var (
	//Generic errors that can be used by any package
	ErrorConvertingStringToBool     = errors.New("The given data isn’t a boolean type value. Provide a valid boolean type in the command’s flag value and try again. Use the flags -h or --help with a command or subcommand to display more information and try again")
	ErrorHandlingFile               = errors.New("The file name doesn’t exist or is invalid. Provide a valid and/or existing path and file name. Use the flags -h or --help with a command or subcommand to display more information and try again")
	ErrorEmptyFile                  = errors.New("The file’s content is empty. Provide a path and file with valid content and try the command again. Use the flags -h or --help with a command or subcommand to display more information and try again")
	ErrorOpeningFile                = errors.New("Failed to open a file. Verify if the path and file exists and/or the file is corrupted and try the command again")
	ErrorInvalidVariablesFileFormat = errors.New("The format of the variables in the file is invalid. You must provide a file with valid variable formats. Use the flags -h or --help with a command or subcommand to display more information")
	ErrorInternalServerError        = errors.New("The server could not process the request because an internal and unexpected problem occurred. Wait a few seconds and try again. Contact Azion’s support for more information")
	ErrorUpdateNoFlagsSent          = errors.New("The subcommand update needs at least one flag with a valid value. Run the command `azioncli <command> update --help` to display more information and try again")
	ErrorUnmarshalReader            = errors.New("Failed to decode the given 'azion.json' file. Verify if the file format is JSON or fix its content according to the JSON format specification at https://www.json.org/json-en.html")
	ErrorFormatOut                  = errors.New("The server failed formatting data for display. Repeat the HTTP request and check the HTTP response's format")
	ErrorWriteFile                  = errors.New("The file is read-only and/or isn't accessible. Change the attributes of the file to read and write and/or give access to it")
	ErrorTokenManager               = errors.New("Internal token handling failure. Run 'azioncli configure --help' command to display more information and try again")
	ErrorTokenNotProvided           = errors.New("Token was not provided; the CLI uses a previous stored token if it was configured. You must provide a valid token, or create a new one, and configure it to use with the CLI. Manage your personal tokens on RTM using the Account Menu > Personal Tokens and configure the token again with the command 'azioncli configure --token <token>'")
	ErrorInvalidToken               = errors.New("The provided token is invalid. You must create a new token and configure it to use with the CLI. Manage your personal tokens on RTM using the Account Menu > Personal Tokens and configure the new token with the command 'azioncli configure --token <new_token>'")
	ErrorToken401                   = errors.New("The token doesn't exist or has expired. Manage your personal tokens on RTM using the Account Menu > Personal Tokens and configure a valid token with the command 'azioncli configure --token <my_token>'")
	ErrorForbidden403               = errors.New("You do not have the permissions to access the API. Make sure the feature is enabled in your profile")
	ErrorNotFound404                = errors.New("The given web page URL or API's endpoint doesn't exist or isn't available. Check that the identifying information is correct. If the error persists, contact Azion's support")
	ErrorFetchingTemplates          = errors.New("Failed to fetch templates from the Azion's GitHub remote repository. Verify the connectivity to the repository https://github.com/aziontech/azioncli-template and try again")
	ErrorMovingFiles                = errors.New("Failed to initialize your project with the Azion template. Please verify if you have write permissions to this directory")
	ErrorUnsupportedType            = errors.New("The project type isn’t supported. Modify the project to a valid type <javascript | nextjs | flareact> and try the command again. Use the flags -h or --help with a command or subcommand to display more information and try again")
	ErrorInvalidOption              = errors.New("You must inform 'yes' or 'no' as input, or force --yes or --no by using the flags")
	ErrorCleaningDirectory          = errors.New("Failed to clean the directory's contents because the directory is read-only and/or isn't accessible. Change the attributes of the directory to read/write and/or give access to it")
	ErrorRunningCommand             = errors.New("Failed to run the command specified in the template (config.json)")

	ErrorLoadingEnvVars         = errors.New("Failed to load the edge applications's environment variables. Verify if the environment variables exist and/or if their values are valid and try again")
	ErrorOpeningAzionJsonFile   = errors.New("Failed to open the given 'azion.json' file. Verify if the file format is JSON or fix its content according to the JSON format specification at https://www.json.org/json-en.html")
	ErrorUnmarshalAzionJsonFile = errors.New("Failed to parse the given 'azion.json' file. Verify if the file format is JSON or fix its content according to the JSON format specification at https://www.json.org/json-en.html")
	ErrorMarshalAzionJsonFile   = errors.New("Failed to encode the given 'azion.json' file. Verify if the file format is JSON or fix its content according to the JSON format specification at https://www.json.org/json-en.html")
	ErrorWritingAzionJsonFile   = errors.New("Failed to write in the given 'azion.json' file. Verify if the file is writable and/or you have access to it, if the data format is JSON, or fix the content according to the JSON format specification at https://www.json.org/json-en.html")
	ErrorTimeoutAPICall         = errors.New("CLI's request has timed out during communication with Azion. Verify if it has completed successfully or wait some time and try the command again")
	ErrorCreateFile             = errors.New("Failed to create %s file")
	ErrorProductNotOwned        = errors.New("This account does not own the following product")
)
