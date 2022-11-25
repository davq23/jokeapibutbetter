<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
    <xsl:output method="html" encoding="utf-8" indent="yes" />
    <xsl:template  match="/">
        <xsl:text disable-output-escaping='yes'>&lt;!DOCTYPE html&gt;</xsl:text>
        <html>
            <head></head>
            <body>
                <p>
                    <xsl:text>Greetings, someone sent you this cracking joke:</xsl:text>
                </p>
                <table>
                    <tbody>
                        <xsl:for-each select="/response/data/joke">
                            <tr>
                                <td>
                                    <xsl:value-of select="text" output-escaping="html" />
                                </td>
                                <td>
                                    <xsl:text>Written by </xsl:text>
                                    <span>
                                        <xsl:value-of select="user/username"/>
                                    </span>
                                </td>
                            </tr>
                        </xsl:for-each>
                    </tbody>
                </table>
            </body>
        </html>
    </xsl:template>

    <xsl:template match="*[contains(name(), '\n')]">
        <br />
    </xsl:template>
</xsl:stylesheet>