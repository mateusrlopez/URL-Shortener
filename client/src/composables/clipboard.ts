export async function copyTextToClipboard(data: string): Promise<void> {
    await navigator.clipboard.writeText(data);
}
