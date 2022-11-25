<?php

class JokeSender
{
    private string $secret;

    public function __construct()
    {
        
    }

    /**
     * Perform XML document based on an XSLT stylesheet
     *
     * @param DOMDocument $xmlDocument
     * @param DOMDocument $xsltDocument
     * @return string|boolean|null
     */
    public function performXSLTTransform(DOMDocument $xmlDocument, DOMDocument $xsltDocument)
    {
        $processor = new XSLTProcessor();

        $processor->importStylesheet($xsltDocument);

        return $processor->transformToXml($xmlDocument);
    }

    public function loadXSLT(string $filename)
    {
        $xsltFileContents = file_get_contents($filename);

        $domDocument = new DOMDocument();

        if (!$domDocument->loadXML($xsltFileContents)) {
            return false;
        }

        return $domDocument;
    }
}

if ($_SERVER["REQUEST_METHOD"] !== "POST") {
    http_response_code(405);
    exit;
}

$jokeSender = new JokeSender();

$jsonBody = json_decode(file_get_contents("php://input"), true);

if (isset($jsonBody["joke_id"])) {
    $xsltDocument = $jokeSender->loadXSLT("transformations/joke_email.xslt");

    if ($xsltDocument === false) {
        http_response_code(400);
        exit;
    }

    $curl = curl_init("http://testapp-19995.nodechef.com/api/jokes/".$jsonBody["joke_id"]."?format=xml");

    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

    $xmlResponse = curl_exec($curl);

    curl_close($curl);

    if ($xsltDocument == false) {
        http_response_code(400);
        exit;
    }

    $xmlDocument = new DOMDocument();
    $xmlDocument->loadXML($xmlResponse);

    $textNodes = $xmlDocument->getElementsByTagName("text");

    /** @var DOMNode $textNode */
    foreach ($textNodes as $textNode) {

    }

    if ($xmlDocument === false) {
        http_response_code(400);
        exit;
    }

    echo $jokeSender->performXSLTTransform($xmlDocument, $xsltDocument);
} else {
    http_response_code(500);
    exit;
}