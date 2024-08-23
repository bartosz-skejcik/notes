<script lang="ts">
	let { children }: { children: any } = $props();
</script>

<div class="grid w-full h-screen grid-cols-5 overflow-hidden">
	{@render children()}
</div>
