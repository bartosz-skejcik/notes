<script lang="ts">
	import { Button } from '$ui/button';
	import { ArrowRight, NotebookText, Trash } from 'lucide-svelte/icons';
	import type { PageData } from '../../routes/$types';

	let { notebooks, session }: PageData = $props();

	function handleClick() {
		window.location.href = '/new-notebook';
	}
</script>

<main
	class="flex flex-col items-center justify-center flex-1 w-full h-screen px-20 text-center gap-y-12"
>
	<div class="flex flex-col items-center space-y-4">
		<NotebookText strokeWidth={1.4} class="p-4 text-foreground bg-muted rounded-xl w-28 h-28" />
		<h1 class="text-3xl xl:text-4xl">{session?.display_name.split(' ')[0]}{"'"}s Notes</h1>
	</div>
	<Button
		class="flex items-center justify-center gap-2 text-foreground bg-brand hover:bg-brand/90"
		variant="default"
		on:click={handleClick}
	>
		<span>Create a new notebook</span>
		<ArrowRight class="w-4 h-4" />
	</Button>
	<div class="flex flex-col items-center w-full max-w-xs space-y-4">
		<p class="text-muted-foreground">or open an existing one</p>
		<div class="flex flex-col items-center justify-center w-full gap-y-1">
			{#if notebooks}
				{#each notebooks! as notebook}
					<div class="flex items-center justify-between w-full gap-2 px-4 py-2 rounded-md">
						<span>{notebook.name}</span>
						<div class="flex items-center gap-2">
							<Button variant="ghost" class="px-3 hover:text-red-500">
								<Trash class="w-4 h-4" />
							</Button>
							<Button class="text-foreground bg-brand hover:bg-brand/90" variant="default"
								>Open</Button
							>
						</div>
					</div>
				{/each}
			{/if}
		</div>
	</div>
	<div class="flex items-center justify-center w-full gap-5">
		<a href="/logout" class="text-muted-foreground hover:text-foreground">Sign out</a>
		<a
			href="https://github.com/bartosz-skejcik/notes"
			class="text-muted-foreground hover:text-foreground">Github</a
		>
	</div>
</main>
