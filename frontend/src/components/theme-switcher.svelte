<script lang="ts">
	import { toggleMode } from 'mode-watcher';
	import { Button } from '$ui/button';

	import { Sun, Moon } from 'lucide-svelte/icons';
</script>

<Button onclick={toggleMode} size="icon" variant="ghost" class={'relative text-muted-foreground'}>
	<Sun
		class="absolute h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all duration-75 dark:-rotate-90 dark:scale-0"
	/>
	<Moon
		class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all duration-75 dark:rotate-0 dark:scale-100"
	/>
	<span class="sr-only">Toggle theme</span>
</Button>
