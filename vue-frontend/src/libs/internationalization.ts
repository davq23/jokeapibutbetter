export const getFlagByLanguage = (languageCode: string): string => {
    switch (languageCode) {
        case 'fr-FR':
            return '🇫🇷';
    }

    return '';
};
