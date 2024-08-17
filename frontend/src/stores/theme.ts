import { writable } from 'svelte/store';
import { applyTheme } from '$lib/themes';
import { themes, DEFAULT_THEME } from '$themes';

const THEME_STORAGE_KEY = 'app-theme';

const getInitialTheme = (): string => {
	if (typeof window !== 'undefined') {
		const storedTheme = localStorage.getItem(THEME_STORAGE_KEY);
		if (storedTheme && themes.hasOwnProperty(storedTheme)) {
			return storedTheme;
		}
	}
	return DEFAULT_THEME;
};

const createThemeStore = () => {
	const { subscribe, set } = writable(getInitialTheme());

	return {
		subscribe,
		set: (value: string) => {
			if (typeof window !== 'undefined') {
				localStorage.setItem(THEME_STORAGE_KEY, value);
			}
			set(value);
		}
	};
};

export const theme = createThemeStore();

theme.subscribe((value) => {
	applyTheme(themes[value]);
});

export const setTheme = (newTheme: string): void => {
	theme.set(newTheme);
};
