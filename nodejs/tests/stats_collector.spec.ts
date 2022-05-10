'use strict';

import { expect } from 'chai';
import { StatsCollector } from '../src/index';

describe('StatsCollector tests', () => {
	it('should create an instance using its constructor', () => {
		const collector: StatsCollector = new StatsCollector();
		expect(collector, 'example should exist').to.exist;
	});
});
