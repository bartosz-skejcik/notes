<script lang="ts">
	import * as Card from '$ui/card';
	import { Input } from '$ui/input';
	import { Label } from '$ui/label';
	import { Button } from '$ui/button';
	import { ArrowLeft } from 'lucide-svelte';

	import { superForm } from 'sveltekit-superforms';

	export let data;

	const { form, errors, enhance, constraints } = superForm(data.form);

	function onCancel() {
		window.location.href = '/';
	}
</script>

<main
	class="flex flex-col items-center justify-center flex-1 w-full h-screen px-20 text-center gap-y-12"
>
	<form method="POST" use:enhance>
		<Card.Root class="w-full max-w-sm bg-transparent border-none shadow-none text-start">
			<Card.Header class="w-full">
				<Card.Title class="w-full text-2xl text-start">Create a new notebook</Card.Title>
			</Card.Header>
			<Card.Content class="grid gap-4">
				<div class="grid gap-2">
					<Label for="name" class="flex flex-col items-start font-semibold text-[0.9em]">
						<p>Name</p>
						<p class="text-sm font-normal text-muted-foreground">Pick a name for your notebook</p>
					</Label>
					<Input
						id="name"
						class="col-span-3"
						name="name"
						type="text"
						placeholder="Personal"
						aria-invalid={$errors.name ? 'true' : undefined}
						bind:value={$form.name}
						{...$constraints.name}
					/>
				</div>
			</Card.Content>
			<Card.Footer class="grid grid-cols-3 gap-2">
				<Button
					variant="secondary"
					class="flex items-center justify-center col-span-1 gap-2"
					onclick={onCancel}
				>
					<ArrowLeft class="w-4 h-4" />
					Back
				</Button>
				<Button type="submit" class="col-span-2">Create</Button>
			</Card.Footer>
		</Card.Root>
	</form>
</main>
