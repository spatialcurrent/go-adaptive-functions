// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

const { concat, split, join, sum } = global.af;

function log(str) {
  console.log(str.replace(/\n/g, "\\n").replace(/\t/g, "\\t").replace(/"/g, "\\\""));
}

describe('af', () => {

  it('split', () => {
    var { result, err } = split("hello world", " ");
    expect(err).toBeNull();
    expect(result).toEqual(["hello", "world"]);
  });

  it('join', () => {
    var { result, err } = join(["hello", "world"], " ");
    expect(err).toBeNull();
    expect(result).toEqual("hello world");
  });

  it('concat', () => {
    var { result, err } = concat(["hello", " ", "world"]);
    expect(err).toBeNull();
    expect(result).toEqual("hello world");
  });

  it('sum', () => {
    var { result, err } = sum([1, [2, 3], 4]);
    expect(err).toBeNull();
    expect(result).toEqual(10);
  });

});
