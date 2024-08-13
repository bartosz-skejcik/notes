export async function load() {
	type Data = {
		notebooks: {
			id: string;
			icon: string;
			name: string;
		}[];
	};

	const d: Data = {
		notebooks: [
			{
				id: '1',
				icon: '📓',
				name: 'Personal'
			},
			{
				id: '2',
				icon: '📣',
				name: 'Social media'
			}
		]
	};

	return d;
}
