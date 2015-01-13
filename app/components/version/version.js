'use strict';

angular.module('tinygraphs.version', [
  'tinygraphs.version.interpolate-filter',
  'tinygraphs.version.version-directive'
])

.value('version', '0.0.1');
