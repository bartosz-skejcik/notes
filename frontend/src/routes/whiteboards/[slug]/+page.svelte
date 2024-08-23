<script lang="ts">
	import ThemeSwitcher from '$components/theme-switcher.svelte';

	import { Button } from '$ui/button';
	import { Home, Plus, Settings, Trash } from 'lucide-svelte';
	import type { PageData } from './$types';
	import Note from '$components/whiteboard/note.svelte';

	let { data }: { data: PageData } = $props();
</script>

<nav class="sticky top-0 flex items-center justify-end px-3 py-2 z-4 bg-background">
	<!-- <div class="flex items-center justify-center gap-4">

	</div> -->
	<div class="flex items-center justify-center gap-4">
		<Button variant="red_ghost" size="sm" class="border border-red-500/20">
			<Trash class="w-4 h-4 mr-2" />
			Delete
		</Button>
		<Button variant="outline" size="sm">
			<Plus class="w-4 h-4 mr-2" />
			Add item
		</Button>
		<div class="flex items-center justify-center gap-3">
			<ThemeSwitcher />
			<Button size="icon" variant="ghost" class={'relative text-muted-foreground w-8 h-8'}>
				<Settings class="w-5 h-5" />
			</Button>
			<Button
				onclick={() => (window.location.href = '/')}
				size="icon"
				variant="ghost"
				class="w-8 h-8"
			>
				<Home class="w-5 h-5 text-muted-foreground" />
			</Button>
		</div>
	</div>
</nav>
<div
	class="grid flex-1 w-full h-full grid-cols-3 gap-3 p-3 pb-20 place-content-start xl:grid-cols-5 lg:grid-cols-4"
>
	{#each data.sticky_notes as sticky_note}
		<Note categories={data.categories} {sticky_note} />
	{/each}
</div>
