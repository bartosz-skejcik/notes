<script lang="ts">
	import { Button } from '$ui/button';
	import { ArrowRight } from 'lucide-svelte/icons';
	import type { PageData } from '../../routes/$types';
	import NotebookButton from '$components/notebook-button.svelte';
	import { useNotebooksStore } from '$stores/notebook.svelte';
	import { onMount } from 'svelte';

	let { session, sessionId }: PageData = $props();

	function handleClick() {
		window.location.href = '/new-notebook';
	}

	let notebookStore = useNotebooksStore();

	let openedNotebook = $state<number | null>(null);

	function toggleNotebook(i: number) {
		if (openedNotebook === i) {
			openedNotebook = null;
		} else {
			openedNotebook = i;
		}
	}

	onMount(async () => {
		if (notebookStore && sessionId) {
			await notebookStore.getNotebooks(sessionId);
		}
	});

	$inspect(notebookStore.notebooks);
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
		{#if notebookStore.notebooks.length > 0}
			<p class="text-muted-foreground">or open an existing one</p>
		{/if}
		<div
			class="relative z-10 flex flex-col items-center justify-center w-full rounded-lg gap-y-0.5 divide-background"
		>
			{#if notebookStore.notebooks}
				{#each notebookStore.notebooks! as notebook, i}
					<NotebookButton {notebook} {i} {toggleNotebook} open={openedNotebook === i} />
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
