import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './language-selector.less!';
import template from './language-selector.stache!';
import Languages from 'gophergala/models/languages';

export const ViewModel = Map.extend({
  define: {
    selectedLanguages: {
      value: []
    },
    languages: {
      value: [],
    },
    getLangs: {
      get() {
        return Languages.findOne({}).then((response)=>{
          this.attr('languages', response.languages.attr());
          this.attr('selectedLanguages', response.languages.attr());
        });
      }
    },
    query: {
      value: ''
    },
    filteredLanguages: {
      value: [],
      get() {
        let query = this.attr('query');
        let matches = []
        let substringRegex = new RegExp(query, 'i');
        if(query) {
          this.attr('languages').each((str)=>{
            if (substringRegex.test(str)) {
              matches.push(str);
            }
          });
          return matches;
        }
        return this.attr('languages');
      }
    }
  },
  removeSelectedLanguage(lang, a, ev) {
    ev.preventDefault();
    let index = this.selectedLanguages.indexOf(lang);
    if(index !== -1) {
      this.selectedLanguages.splice(index, 1);
    }
  },
  selectLanguage(lang, a, ev) {
    ev.preventDefault();
    ev.stopPropagation();
    if(this.attr('selectedLanguages').indexOf(lang) == -1) {
      this.attr('selectedLanguages').push(lang);
    }
    this.attr('query', '');
  },
  setQuery(map, el) {
    this.attr('query', el.val());
  },
  newEntry(map, el, ev) {
    ev.preventDefault();
    ev.stopPropagation();
    this.attr('selectedLanguages').push(el.val());
    el.val('');
  }
});

let clickHandler = function () {
   $('.filtered-langs').css('display', 'none');
}

export default Component.extend({
  tag: 'language-selector',
  viewModel: ViewModel,
  template,
  events: {
    'input.type-ahead click': function (el, ev) {
      ev.stopPropagation();
      $('.filtered-langs').css('display', 'block');
      $('body').click(clickHandler);
    },
    'removed': function  (argument) {
      $('body').off(clickHandler);
    }
  }
});