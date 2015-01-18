'use strict';

describe('tinygraphs.version module', function() {
  beforeEach(module('tinygraphs.version'));

  describe('version service', function() {
    it('should return current version', inject(function(version) {
      expect(version).toEqual('0.0.3 Running Free');
    }));
  });
});
