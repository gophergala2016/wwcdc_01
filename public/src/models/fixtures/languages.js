import fixture from 'can-fixture';

let findAllLanguages = function () {
  return {
    languages: [
      "Algorithms",
      "C",
      "C++",
      "Java",
      "More",
      "Python"
    ]
  }
}

fixture({
  'GET api/languages': findAllLanguages,
});
