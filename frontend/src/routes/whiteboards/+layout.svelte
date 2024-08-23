<script lang="ts">
	import Sidebar from '$components/whiteboard/sidebar.svelte';
	import type { LayoutData } from './$types';

	const { data, children }: { data: LayoutData; children: any } = $props();

	console.log(data);
</script>

<div class="grid w-full h-screen grid-cols-5 overflow-hidden">
	<aside class="relative col-span-1 bg-background">
		<Sidebar categories={data.categories} view={data.view} />
	</aside>
	<main class="col-span-4 overflow-auto">
		{@render children(data)}
	</main>
</div>
