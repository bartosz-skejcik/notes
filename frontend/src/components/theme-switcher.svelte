<script lang="ts">
	import { Button } from '$ui/button';

	import { Sun, Moon, StickyNote } from 'lucide-svelte/icons';
	import { theme, setTheme } from '$stores/theme';
	import { onMount } from 'svelte';

	onMount(() => {
		setTheme($theme);
	});

	function cycleTheme(): void {
		const themes = ['base', 'dark', 'paper'];
		const currentIndex = themes.indexOf($theme);
		const nextIndex = (currentIndex + 1) % themes.length;
		setTheme(themes[nextIndex]);
	}
</script>

<Button onclick={cycleTheme} size="icon" variant="ghost" class={'relative text-muted-foreground'}>
	{#if $theme === 'base'}
		<Sun class="absolute h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all duration-75" />
	{/if}
	{#if $theme === 'dark'}
		<Moon class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all duration-75" />
	{/if}
	{#if $theme === 'paper'}
		<StickyNote
			class="absolute h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all duration-75"
		/>
	{/if}
	<span class="sr-only">Toggle theme</span>
</Button>
