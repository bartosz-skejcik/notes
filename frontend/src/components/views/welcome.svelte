<script lang="ts">
	import { Button } from '$ui/button';
	import { ArrowRight, Trash } from 'lucide-svelte/icons';
	import type { PageData } from '../../routes/$types';

	let { notebooks, session }: PageData = $props();

	function handleClick() {
		window.location.href = '/new-notebook';
	}

	function navigate(path: string) {
		window.location.href = path;
	}
</script>

<main
	class="flex flex-col items-center justify-center flex-1 w-full h-screen px-20 text-center gap-y-12"
>
	<div class="flex flex-col items-center space-y-4 text-foreground">
		<!-- <NotebookText strokeWidth={1.4} class="p-4 text-foreground bg-muted rounded-xl w-28 h-28" /> -->

		<img src="/logo.png" alt="logo" class="w-36 h-36" />

		<h1 class="text-3xl xl:text-4xl">{session?.display_name.split(' ')[0]}{"'"}s Notes</h1>
	</div>
	<Button class="flex items-center justify-center gap-2" variant="default" on:click={handleClick}>
		<span>Create a new notebook</span>
		<ArrowRight class="w-4 h-4" />
	</Button>
	<div class="flex flex-col items-center w-full max-w-xs space-y-4">
		{#if notebooks.length > 0}
			<p class="text-muted-foreground">or open an existing one</p>
		{/if}
		<div
			class="flex flex-col items-center justify-center w-full divide-y rounded-lg gap-y-1 bg-foreground/10 divide-background"
		>
			{#if notebooks}
				{#each notebooks! as notebook}
					<div class="flex items-center justify-between w-full gap-2 py-1 pl-3 pr-2">
						<div class="flex flex-col items-start justify-center">
							<span class="font-medium">{notebook.name}</span>
							<span class="text-xs text-muted-foreground"
								>{new Date(notebook.created_at).toLocaleDateString('en-PL', {
									year: 'numeric',
									month: 'long',
									day: 'numeric'
								})}</span
							>
						</div>
						<div class="flex items-center gap-2">
							<Button variant="red_ghost" class="px-3">
								<Trash class="w-4 h-4" />
							</Button>
							<Button
								on:click={() => navigate(`/notebooks/${notebook.id}`)}
								variant="ghost"
								class="bg-foreground/10 hover:bg-foreground/20">Open</Button
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
