<script lang="ts">
	import type { Category } from '$lib/whiteboard/categories';
	import { Plus } from 'lucide-svelte';
	import { Button } from '$ui/button';

	type Props = {
		categories: Category[];
		view: string;
	};

	let { categories, view }: Props = $props();

	let currentCategory = $derived(categories.find((c) => c.id.toString() === view) ?? null);
</script>

<div
	class="flex flex-col items-start justify-start w-full h-full py-2 pl-4 pr-1 border-r border-foreground/10"
>
	<header class="flex items-start justify-start w-full gap-2 py-1 mb-4 text-sm text-foreground/80">
		<span class="whitespace-nowrap">My Library</span>
		<span class="text-muted-foreground">{'/'}</span>
		<span class="whitespace-nowrap">{currentCategory?.name ?? 'Recents'}</span>
	</header>
	<div class="flex flex-col items-start justify-start w-full mb-2 gap-y-1">
		<a
			href="?view=recent"
			class={`flex items-center justify-start gap-1 py-1 text-sm text-foreground ${view === 'recent' ? 'font-medium' : 'font-normal text-foreground/70'}`}
		>
			Recents
		</a>
		<a
			href="?view=all"
			class={`flex items-center justify-start gap-1 py-1 text-sm text-foreground ${view === 'all' ? 'font-medium' : 'font-normal text-foreground/70'}`}
		>
			All
		</a>
	</div>
	<div class="flex flex-col items-start justify-start flex-1 w-full h-full mt-3 gap-y-1">
		<div class="flex items-center justify-between w-full">
			<p class="flex items-center justify-start gap-1 py-1 text-sm font-medium text-foreground">
				My Library
			</p>
			<!-- small add category button -->
			<Button size="icon" variant="ghost" class="w-6 h-6 lg:hidden">
				<Plus class="w-4 h-4" />
			</Button>
			<!-- big add category button -->
			<a href="?open=true">
				<Button variant="ghost" class="px-2 text-xs h-7">
					<Plus class="w-4 h-4 mr-2" />
					Add
				</Button>
			</a>
		</div>
		{#each categories as category}
			<a
				href="?view={category.id}"
				class={`flex items-center justify-start gap-2 py-1 text-sm ${
					category.id.toString() === view
						? 'font-medium text-foreground'
						: 'font-normal text-foreground/70'
				}`}
			>
				<div style={`background: ${category.color}`} class={`rounded-full w-3.5 h-3.5`}></div>
				{category.name}
			</a>
		{/each}
	</div>
</div>
