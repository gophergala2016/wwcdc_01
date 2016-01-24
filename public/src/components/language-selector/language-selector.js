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
      value: '',
      set(val, setter) {
        setter(val);
        this.attr('filteredLanguages');
      }
    },
    filteredLanguages: {
      value: [],
      get() {
        let query = this.attr('query');
        let matches = []
        let substringRegex = new RegExp(query, 'i');
        if(query) {
          this.attr('languages').each(strs, function(i, str) {
            if (substrRegex.test(str)) {
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
    if(!this.attr('selectedLanguages').indexOf(lang)) {
      this.attr('selectedLanguages').push(lang);
    }
    this.attr('query', '');
  }
});

export default Component.extend({
  tag: 'language-selector',
  viewModel: ViewModel,
  template
});