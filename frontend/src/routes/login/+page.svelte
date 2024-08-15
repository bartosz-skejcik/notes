<script lang="ts">
	import { Button } from '$ui/button';
	import * as Card from '$ui/card';
	import { Input } from '$ui/input';
	import { Label } from '$ui/label';

	import { superForm } from 'sveltekit-superforms';

	export let data;

	const { form, errors, enhance, constraints } = superForm(data.form);

	$: console.log(data);
</script>

<main
	class="flex flex-col items-center justify-center flex-1 w-full h-screen px-20 text-center gap-y-12"
>
	<form method="POST" use:enhance>
		<Card.Root class="w-full max-w-sm text-start">
			<Card.Header>
				<Card.Title class="text-2xl">Login</Card.Title>
				<Card.Description>Enter your email below to login to your account.</Card.Description>
			</Card.Header>
			<Card.Content class="grid gap-4">
				<div class="grid gap-2">
					<Label for="email">Email</Label>
					<Input
						type="email"
						name="email"
						placeholder="m@example.com"
						aria-invalid={$errors.email ? 'true' : undefined}
						bind:value={$form.email}
						{...$constraints.email}
					/>
					{#if $errors.email}
						<span class="text-sm font-normal text-red-500">{$errors.email}</span>
					{/if}
				</div>
				<div class="grid gap-2">
					<Label for="password">Password</Label>
					<Input
						type="password"
						name="password"
						aria-invalid={$errors.password ? 'true' : undefined}
						bind:value={$form.password}
						{...$constraints.password}
					/>
					{#if $errors.password}
						<span class="text-sm font-normal text-red-500">{$errors.password}</span>
					{/if}
				</div>
			</Card.Content>
			<Card.Footer class="flex flex-col items-start justify-center w-full gap-2 text-start">
				{#if data.error && data.error.length > 0}
					<p class="text-red-500">{data.error}</p>
				{/if}
				<Button type="submit" class="w-full">Sign in</Button>
			</Card.Footer>
		</Card.Root>
	</form>
</main>
