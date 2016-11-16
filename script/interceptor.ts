class XHRInterceptor {
  private static readonly xOpen = XMLHttpRequest.prototype.open;
  private static readonly xSend = XMLHttpRequest.prototype.send;
  private static openOR = false;
  private static sendOR = false;

  private matchers: Matcher[] = [];

  constructor(private targetURL: string = 'http://localhost:18000/') { }

  match(matcher: Matcher) {
    this.matchers.push(matcher);
  }

  matchAll(matchers: Matcher[]) {
    this.matchers.push(...matchers);
  }

  restore() {
    XMLHttpRequest.prototype.open = XHRInterceptor.xOpen;
    XMLHttpRequest.prototype.send = XHRInterceptor.xSend;
  }

  static restore() {
    XMLHttpRequest.prototype.open = XHRInterceptor.xOpen;
    XMLHttpRequest.prototype.send = XHRInterceptor.xSend;
  }

  intercept() {
    this.interceptOpen(XMLHttpRequest.prototype.open);
    this.interceptSend(XMLHttpRequest.prototype.send);
  }

  private interceptOpen(open: Function) {
    if (XHRInterceptor.openOR) {
      return;
    }

    const self = this;

    XMLHttpRequest.prototype.open = function (method: string, url: string) {
      this.originalURL = url;
      method = 'POST';

      open.call(this, method, self.targetURL);
    };
  }

  private interceptSend(send: Function) {
    if (XHRInterceptor.sendOR) {
      return;
    }

    const self = this;

    XMLHttpRequest.prototype.send = function (data: any) {

      const parsedURL = new URL(this.originalURL);

      const l = self.matchers.filter((m) => {
        if (typeof m.expression == 'string') {
          m.expression = new RegExp(m.expression);
        }

        const matches = parsedURL.pathname.match(m.expression);


        return (
          matches &&
          matches.length > 0 &&
          matches[0].length == parsedURL.pathname.length
        );
      });

      const payload = { response: { result: 'default payload' }, statusCode: 200 };

      if (l.length > 0) {
        payload.statusCode = l[0].statusCode || 200;
        payload.response = l[0].response || { result: 'default payload' };
      } else {
        const consoleArgs: any[] = [];
        let message = `%cURL not handled: %c${this.originalURL}`;
        if (data) {
          message += `\n%cData: %c`;
          message += typeof data == 'string' ? `%s` : `%O`;
          consoleArgs.push(
            `font-style: italic; color: red; font-size: 12px`,
            `font-style: normal; color: green; font-size: 12px;`,
            data,
          );
        }
        consoleArgs.unshift(
          `font-style: italic; color: red; font-size: 12px`,
          `font-style: normal; color: green; font-size: 12px;`,
        );
        console.log(message, ...consoleArgs);
      }

      if (typeof payload == 'object') {

        this.setRequestHeader('Content-Type', 'application/json');
        send.call(this, JSON.stringify(payload));

      } else {
        send.call(this, payload);
      }
    };
  }
}

(window as any).XHRInterceptor = XHRInterceptor;

class Matcher {
  constructor(
    public expression: RegExp | string,
    public response: any,
    public statusCode: number,
  ) { }
}
