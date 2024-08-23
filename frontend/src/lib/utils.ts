import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
import { cubicOut } from 'svelte/easing';
import type { TransitionConfig } from 'svelte/transition';

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

type FlyAndScaleParams = {
	y?: number;
	x?: number;
	start?: number;
	duration?: number;
};

export const flyAndScale = (
	node: Element,
	params: FlyAndScaleParams = { y: -8, x: 0, start: 0.95, duration: 150 }
): TransitionConfig => {
	const style = getComputedStyle(node);
	const transform = style.transform === 'none' ? '' : style.transform;

	const scaleConversion = (valueA: number, scaleA: [number, number], scaleB: [number, number]) => {
		const [minA, maxA] = scaleA;
		const [minB, maxB] = scaleB;

		const percentage = (valueA - minA) / (maxA - minA);
		const valueB = percentage * (maxB - minB) + minB;

		return valueB;
	};

	const styleToString = (style: Record<string, number | string | undefined>): string => {
		return Object.keys(style).reduce((str, key) => {
			if (style[key] === undefined) return str;
			return str + `${key}:${style[key]};`;
		}, '');
	};

	return {
		duration: params.duration ?? 200,
		delay: 0,
		css: (t) => {
			const y = scaleConversion(t, [0, 1], [params.y ?? 5, 0]);
			const x = scaleConversion(t, [0, 1], [params.x ?? 0, 0]);
			const scale = scaleConversion(t, [0, 1], [params.start ?? 0.95, 1]);

			return styleToString({
				transform: `${transform} translate3d(${x}px, ${y}px, 0) scale(${scale})`,
				opacity: t
			});
		},
		easing: cubicOut
	};
};

export const colors = [
	'#94a3b8',
	'#64748b',
	'#9ca3af',
	'#4b5563',
	'#a1a1aa',
	'#71717a',
	'#a3a3a3',
	'#737373',
	'#a8a29e',
	'#78716c',
	'#f87171',
	'#ef4444',
	'#fb923c',
	'#f97316',
	'#fbbf24',
	'#f59e0b',
	'#facc15',
	'#eab308',
	'#a3e635',
	'#84cc16',
	'#4ade80',
	'#22c55e',
	'#34d399',
	'#10b981',
	'#2dd4bf',
	'#14b8a6',
	'#22d3ee',
	'#06b6d4',
	'#38bdf8',
	'#0ea5e9',
	'#60a5fa',
	'#3b82f6',
	'#818cf8',
	'#6366f1',
	'#a78bfa',
	'#8b5cf6',
	'#c084fc',
	'#a855f7',
	'#e879f9',
	'#d946ef',
	'#f472b6',
	'#ec4899',
	'#fb7185',
	'#f43f5e'
];
