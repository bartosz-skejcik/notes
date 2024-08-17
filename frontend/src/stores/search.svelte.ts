export const createSearch = () => {
	let search = $state('');

	return {
		setValue(value: string) {
			search = value;
		},
		get value() {
			return search;
		}
	};
};
