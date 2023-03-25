<?php declare(strict_types=1);

function handleError(?array $lastError = null) {
    if (!$lastError  || !array_key_exists("message", $lastError)) {
        $lastError = [];
        $lastError["message"] = "Unknown error";
    }

    throw new Exception(sprintf('Unable to connect to FTP: %s', $lastError["message"]));
}

$lastError = null;
$username = getenv('FTP_USERNAME') ?? '';
$password = getenv('FTP_PASSWORD') ?? '';
$hostname = getenv('FTP_HOST') ?? '';
$port = intval(getenv('FTP_PORT') ?? '21');
$timeout = intval(getenv('FTP_TIMEOUT') ?? '90');
$onError = false;

error_log(print_r([
    "USER: $username",
    "HOST: $hostname",
], true));

$ftp = ftp_connect($hostname, $port, $timeout);

if ($ftp === false) {
    handleError(error_get_last());
}

try {
    $ok = ftp_login($ftp, $username, $password);
    
    if (!$ok) {
        handleError(error_get_last());
    }

    ftp_pasv($ftp, true);

    ftp_mkdir($ftp, 'assets');

    @ftp_put(
        $ftp,
        ".htaccess",
        ".htaccess",
        FTP_BINARY
    );
    
    $baseFilesToUpload = glob('dist/*');
    $assetFilesToUpload = glob('dist/assets/*');

    $filesToUpload = array_merge($baseFilesToUpload, $assetFilesToUpload);
    $filesUploaded = [];

    foreach ($filesToUpload as $fileToUpload) {
        if (!str_contains($fileToUpload, ".")) {
            continue;
        }
        $baseFilename = basename($fileToUpload);
        $ok = @ftp_put(
            $ftp,
            (str_contains($fileToUpload, 'assets') ? 'assets/' : '').$baseFilename,
            $fileToUpload,
            FTP_BINARY
        );

        if ($ok) {
            $filesUploaded[] = $fileToUpload;
        } else {
            echo 'Failed to upload: '.__DIR__.DIRECTORY_SEPARATOR.$fileToUpload;
            exit(-1);
        }
    }

    echo 'Files uploaded: '.print_r($filesUploaded, true);
    echo 'Files not uploaded: '.print_r(array_diff($filesToUpload, $filesUploaded), true);
} catch (\Throwable $th) {
    echo 'Error: '.$th->getMessage();
    $onError = true;
} finally {
    ftp_close($ftp);
}

exit($onError ? -1 : 0);