import {expect} from 'chai';
import {List} from 'immutable';

describe('immutability', () => {
  describe('a number', () => {
    function incremet(currentState) {
      return currentState + 1;
    }

    it('is immutable', () => {
      let state = 42;
      let nextState = incremet(state);

      expect(nextState).to.equal(43);
      expect(state).to.equal(42);
    });
  });

  describe('A List', () => {
    function addMovie(currentState, movie) {
      return currentState.push(movie);
    }

    it('is immutable', () => {
      let state = List.of('Trainspotting', '28 Days Later');
      let nextState = addMovie(state, 'Sunshine');

      expect(nextState).to.equal(
        List.of('Trainspotting', '28 Days Later', 'Sunshine'),
      );

      expect(state).to.equal(List.of('Trainspotting', '28 Days Later'));
    });
  });
});
