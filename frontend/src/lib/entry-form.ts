export function autoResize() {
	const textarea = document.querySelector('textarea');
	if (textarea) {
		textarea.style.height = 'auto';
		textarea.style.height = textarea.scrollHeight + 'px';
	}
}

export function handleFileUpload(
	event: Event,
	setFile: (file: File | null) => void,
	setImagePreview: (imagePrev: string | null) => void,
	imagePreview: string | null,
	setImagePreviewDims: (imagePrevDims: { width: number; height: number } | null) => void
) {
	const inputFile = (event.target as HTMLInputElement).files?.[0];
	if (inputFile) {
		setFile(inputFile);
		const reader = new FileReader();
		reader.onload = (e) => {
			setImagePreview(e.target?.result as string);

			const image = new Image();
			image.src = imagePreview!;
			image.onload = () => {
				setImagePreviewDims({
					width: image.width,
					height: image.height
				});
			};
		};
		reader.readAsDataURL(inputFile);
	}
}

export function deleteImage(
	setFile: (file: File | null) => void,
	setImagePreview: (imagePrev: string | null) => void
) {
	setFile(null);
	setImagePreview(null);
	// Reset the file input
	const fileInput = document.getElementById('image-upload') as HTMLInputElement;
	if (fileInput) {
		fileInput.value = '';
	}
}
