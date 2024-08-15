<script lang="ts">
	import { Button } from '$ui/button';
	import { ArrowRight, NotebookText, Trash } from 'lucide-svelte/icons';
	import type { PageData } from '../../routes/$types';

	let { notebooks, session }: PageData = $props();
</script>

<main
	class="flex flex-col items-center justify-center flex-1 w-full h-screen px-20 text-center gap-y-12"
>
	<!-- <div class="w-28 aspect-square bg-muted-foreground"></div> -->
	<NotebookText strokeWidth={1.4} class="p-4 text-white bg-muted rounded-xl w-28 h-28" />
	<!-- TODO: add a logo -->
	<h1 class="text-3xl xl:text-4xl">{session?.display_name}{"'"}s Notes</h1>
	<Button
		class="flex items-center justify-center gap-2 text-white bg-brand hover:bg-brand/90"
		variant="default"
	>
		<span>Create a new notebook</span>
		<ArrowRight class="w-4 h-4" />
	</Button>
	<p class="text-muted-foreground">or open an existing one</p>
	<div class="flex flex-col items-center justify-center w-full max-w-xs gap-y-1">
		{#if notebooks}
			{#each notebooks! as notebook}
				<div
					class="flex items-center justify-between w-full gap-2 px-4 py-2 transition-all duration-200 rounded-md hover:bg-muted"
				>
					<span>{notebook.icon}{' '}{notebook.name}</span>
					<div class="flex items-center gap-2">
						<Button variant="ghost" class="px-3 hover:text-red-500">
							<Trash class="w-4 h-4" />
						</Button>
						<Button class="text-white bg-brand hover:bg-brand/90" variant="default">Open</Button>
					</div>
				</div>
			{/each}
		{/if}
	</div>
</main>
