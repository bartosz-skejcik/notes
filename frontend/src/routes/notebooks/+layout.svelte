<!-- file: src/routes/notebooks/+layout.svelte -->

<script lang="ts">
	import Sidebar from '$components/notebooks/sidebar.svelte';

	const { data, children }: { data: any; children: any } = $props();
</script>

<div class="grid w-full h-screen grid-cols-5 overflow-hidden">
	<aside class="relative col-span-1 bg-muted">
		<Sidebar {...data} />
	</aside>
	<main class="col-span-4 overflow-auto">
		{@render children()}
	</main>
</div>
