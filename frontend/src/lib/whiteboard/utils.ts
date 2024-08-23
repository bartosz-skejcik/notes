export function hexToRGBA(hex: string, alpha = 1) {
    // Remove the hash at the start if it's there
    hex = hex.replace(/^#/, '');

    // Parse the hex values
    let r, g, b;

    if (hex.length === 3) {
        // For shorthand hex (#RGB)
        r = parseInt(hex.charAt(0) + hex.charAt(0), 16);
        g = parseInt(hex.charAt(1) + hex.charAt(1), 16);
        b = parseInt(hex.charAt(2) + hex.charAt(2), 16);
    } else if (hex.length === 6) {
        // For full hex (#RRGGBB)
        r = parseInt(hex.substring(0, 2), 16);
        g = parseInt(hex.substring(2, 4), 16);
        b = parseInt(hex.substring(4, 6), 16);
    } else {
        throw new Error('Invalid hex color format');
    }

    // Ensure alpha is between 0 and 1
    alpha = Math.min(1, Math.max(0, alpha));

    // Return the RGBA string
    return `rgba(${r}, ${g}, ${b}, ${alpha})`;
}