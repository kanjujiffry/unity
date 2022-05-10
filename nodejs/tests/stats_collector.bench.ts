'use strict';

import { expect } from 'chai';
import { StatsCollector } from '../src/index';

const iterations = 10000000;
const requestTimeout = 19000;

describe('StatsCollector benchmark', () => {
	const randomValue = () => Math.floor(Math.random() * requestTimeout)

	it('Benchmark median', () => {
		const collector: StatsCollector = new StatsCollector();
		expect(collector, 'example should exist').to.exist;

		console.time('Start');
		for (let i = 0; i < iterations; i++) {
			collector.record(randomValue());
			collector.median();
		};
		console.timeEnd('End');
	});

	it('Benchmark average', () => {
		const collector: StatsCollector = new StatsCollector();
		expect(collector, 'example should exist').to.exist;

		const iterations = 10000000;
		console.time('Start');
		for (let i = 0; i < iterations; i++) {
			collector.record(randomValue());
			collector.average();
		};
		console.timeEnd('End');
	});
});
