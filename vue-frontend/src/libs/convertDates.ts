export const formatDate = (date: string): string => {
    const [dateString, timeString] = date.split(/[T|.]/g, 3);

    return `${dateString} ${timeString}`;
};
