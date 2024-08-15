export function createSearch() {
	let value = $state('');

	return {
		get value() {
			return value;
		},
		set value(v) {
			value = v;
		},
		setValue(v: string) {
			value = v;
		}
	};
}
