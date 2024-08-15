<script lang="ts">
	let { children } = $props();
</script>

<div class="grid w-full h-screen grid-cols-5 overflow-hidden">
	<aside class="col-span-1 bg-muted">Sidebar</aside>
	<main class="col-span-4 overflow-auto">
		{@render children()}
	</main>
</div>
