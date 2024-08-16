<script lang="ts">
	import { Button } from '$ui/button';
	import * as Card from '$ui/card';
	import { Input } from '$ui/input';
	import { Label } from '$ui/label';

	import { superForm } from 'sveltekit-superforms';

	export let data;

	const { form, errors, enhance, constraints } = superForm(data.form);
</script>

<main
	class="flex flex-col items-center justify-center flex-1 w-full h-screen px-20 text-center gap-y-12"
>
	<form method="POST" use:enhance>
		<Card.Root class="max-w-sm mx-auto border-none text-start">
			<Card.Header>
				<Card.Title class="text-xl">Sign Up</Card.Title>
				<Card.Description>Enter your information to create an account</Card.Description>
			</Card.Header>
			<Card.Content>
				<div class="grid gap-4">
					<div class="grid grid-cols-2 gap-4">
						<div class="grid gap-2">
							<Label for="first-name">First name</Label>
							<Input id="first-name" name="firstName" placeholder="Max" required />
						</div>
						<div class="grid gap-2">
							<Label for="last-name">Last name</Label>
							<Input id="last-name" name="lastName" placeholder="Robinson" required />
						</div>
					</div>
					<div class="grid gap-2">
						<Label for="email">Email</Label>
						<Input id="email" type="email" name="email" placeholder="m@example.com" required />
					</div>
					<div class="grid gap-2">
						<Label for="password">Password</Label>
						<Input id="password" type="password" name="password" />
					</div>
					<Button type="submit" class="w-full">Create an account</Button>
				</div>
				<div class="mt-4 text-sm text-start">
					Already have an account?
					<a href="/login" class="underline"> Sign in </a>
				</div>
			</Card.Content>
		</Card.Root>
	</form>
</main>
