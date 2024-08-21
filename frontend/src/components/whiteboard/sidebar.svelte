<script lang="ts">
	import type { Category } from '$lib/whiteboard/categories';

	type Props = {
		categories: Category[];
	};

	let { categories } = $props();
</script>
