<script lang="ts">
	// import Sidebar from '$components/sidebar.svelte';

	const { children }: { children: any } = $props();
</script>

<div class="grid w-full h-screen grid-cols-5 overflow-hidden">
	<aside class="relative col-span-1 bg-muted">
		<!-- <Sidebar entries={data.entries} /> -->
	</aside>
	<main class="col-span-4 overflow-auto">
		{@render children()}
	</main>
</div>
