<script lang="ts">
	import { Input } from '$ui/input';
	import { Search } from 'lucide-svelte';

	type Props = {
		onChange: (value: string) => void;
		items: {
			label: string;
			value: string;
		}[];
	};
	let { onChange, items }: Props = $props();

	let search = $state('');
	let list = $state(items);
	let resultDiv = $state<HTMLDivElement>();

	let rootDiv: HTMLDivElement;

	$effect(() => {
		onChange(search);
	});

	$effect(() => {
		if (search != '') {
			list = items.filter((item) => item.label.toLowerCase().includes(search.toLowerCase()));
		} else {
			list = items;
		}
	});
</script>

<div
	bind:this={rootDiv}
	class={`relative w-fit border border-border ${search !== '' ? 'rounded-t-lg' : 'rounded-lg '}`}
>
	<div class="flex items-center justify-center gap-2 px-3">
		<Search class="w-4 h-4 text-muted-foreground" />
		<Input
			placeholder="Search..."
			bind:value={search}
			class="!outline-none border-none focus-visible:ring-0 p-0 w-56"
		/>
	</div>
	<div
		style={`top: calc(-1 * ${resultDiv?.offsetHeight}); width: ${rootDiv?.offsetWidth}px`}
		bind:this={resultDiv}
		class={`flex flex-col bg-background px-1 border border-border rounded-b-lg absolute z-20 py-2 -ml-[1px] w-full gap-1 ${search ? '' : 'hidden'}`}
	>
		{#each list as item}
			<div class="px-2 text-muted-foreground">{item.label}</div>
		{/each}
		{#if list.length === 0}
			<div class="text-muted-foreground">No results</div>
		{/if}
	</div>
</div>
