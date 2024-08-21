<script lang="ts">
	import ThemeSwitcher from '$components/theme-switcher.svelte';

	import { Button } from '$ui/button';
	import { Home, Settings } from 'lucide-svelte';
	import type { LayoutData } from '../$types';

	let { data }: { data: LayoutData } = $props();

	console.log(data);
</script>

<nav class="sticky top-0 flex items-center justify-between px-3 py-2 z-4 bg-background">
	<div class="flex items-center gap-2 text-[0.97rem] text-muted-foreground">
		<!-- <p>{data.notebook.name}</p> -->
		<span>•</span>
		<p>
			{new Date().toLocaleString('en-PL', {
				weekday: 'long',
				year: 'numeric',
				month: 'long',
				day: 'numeric'
			})}
		</p>
	</div>
	<div class="flex items-center justify-end w-1/2 gap-4">
		<!-- <CommandSearch items={convertEntriesToItems(entries)} onChange={search.setValue} /> -->
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
<div class="p-3 pb-20 space-y-1">
	<!-- <EntryForm {...data} bind:entries />
	{#each entries as entry, i (i)}
		<div animate:flip={{ duration: 300, easing: quintOut }} class="w-full">
			<Entry bind:entries {entry} {...data} tags={data.tags} />
		</div>
	{/each}
	{#if entries.length === 0}
		<Encuragement />
	{/if} -->
</div>
