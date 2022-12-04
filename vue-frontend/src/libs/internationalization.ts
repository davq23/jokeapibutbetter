export const getFlagClassByLanguage = (
    languageCode: string,
    squared: boolean = false,
): string => {
    languageCode = languageCode.replace('_', '-');

    switch (languageCode) {
        case 'en-US':
        case 'en-UK':
            languageCode = 'gb-US';
            break;

        case 'fr-CA':
            languageCode = 'ca-CA';
            break;

        default:
            break;
    }

    return `fib fi-${languageCode.split('-')[0]} ${squared ? 'fis' : ''}`;
};
