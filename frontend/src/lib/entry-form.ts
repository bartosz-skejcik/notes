export function autoResize() {
	const textarea = document.querySelector('textarea');
	if (textarea) {
		textarea.style.height = 'auto';
		textarea.style.height = textarea.scrollHeight + 'px';
	}
}

export function handleFileUpload(
	files: FileList | null,
	setImage: (files: File) => void,
	setImagePreview: (imagePrev: string | null) => void,
	imagePreview: string | null,
	setImagePreviewDims: (imagePrevDims: { width: number; height: number } | null) => void
) {
	const inputFile = files ? files[0] : undefined;
	if (inputFile) {
		setImage(inputFile);
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
	setImage: (filef: File | undefined) => void,
	setImagePreview: (imagePrev: string | null) => void
) {
	setImage(undefined);
	setImagePreview(null);
	// Reset the file input
	const fileInput = document.getElementById('image_upload') as HTMLInputElement;
	if (fileInput) {
		fileInput.value = '';
	}
}
