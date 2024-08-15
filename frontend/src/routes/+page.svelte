<script lang="ts">
	import Welcome from '$components/views/welcome.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<Welcome {...data} />
