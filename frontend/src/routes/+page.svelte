<script lang="ts">
	import Welcome from '$components/views/welcome.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<Welcome {...data} />
