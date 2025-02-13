<script lang="ts">
	import { ChevronDown, NotebookTabs, StickyNote, Trash } from 'lucide-svelte/icons';
	import { Button } from '$ui/button';
	import type { Notebook } from '../routes/+page.server';
	import { slide } from 'svelte/transition';
	import { cubicInOut } from 'svelte/easing';
	import { useNotebooksStore } from '$stores/notebook.svelte';

	type Props = {
		notebook: Notebook;
		i: number;
		open: boolean;
		toggleNotebook: (i: number) => void;
	};

	let { notebook, i, open, toggleNotebook }: Props = $props();

	function navigate(path: string) {
		window.location.href = path;
	}

	let notebookStore = useNotebooksStore();
</script>

<button
	onclick={() => toggleNotebook(i)}
	class={`relative flex-col flex items-center justify-center w-full gap-2 pt-1 pb-2 pl-3 pr-2 bg-accent opacity-80 hover:opacity-100 group ${i == notebookStore.notebooks.length - 1 ? 'rounded-b-lg' : i == 0 ? 'rounded-t-lg' : ''}`}
>
	<div class="flex items-center justify-between w-full gap-2">
		<div class="relative z-20 flex flex-col items-start justify-center">
			<span class="font-medium">{notebook.name}</span>
			<span class="text-xs text-muted-foreground">
				{new Date(notebook.created_at).toLocaleDateString('en-PL', {
					year: 'numeric',
					month: 'long',
					day: 'numeric'
				})}
			</span>
		</div>
		<div class="relative z-20 flex items-center gap-2">
			<Button variant="red_ghost" class="px-3">
				<Trash class="w-4 h-4" />
			</Button>
			<Button
				variant="ghost"
				size="icon"
				class="w-7 h-7 outline outline-1 outline-muted-foreground hover:bg-muted-foreground/20"
			>
				<ChevronDown class={`w-4 h-4 transition-all duration-200 ${open ? 'rotate-180' : ''}`} />
			</Button>
		</div>
	</div>
	{#if open}
		<div
			transition:slide={{ duration: 100, axis: 'y', easing: cubicInOut }}
			class="flex items-center justify-start w-full gap-2"
		>
			<Button
				onclick={() => navigate(`/notebooks/${notebook.id}`)}
				variant="ghost"
				class="w-1/2 bg-foreground/10 hover:bg-foreground/20"
			>
				<NotebookTabs class="w-4 h-4 mr-2" />
				Journal
			</Button>
			<Button
				onclick={() => navigate(`/whiteboards/${notebook.id}`)}
				variant="ghost"
				class="w-1/2 bg-foreground/10 hover:bg-foreground/20"
			>
				<StickyNote class="w-4 h-4 mr-2" />
				Notes
			</Button>
		</div>
	{/if}
</button>
