export function SetTokenToCookie(token, age) {
  document.cookie = `testbed-token=${token}; max-age=${age}`;
}

export function GetTokenMetadata() {
  const token = GetTokenFromCookie();
  return {'x-testbed-token': token};
}

export function GetTokenFromCookie() {
  const cookies = getCookies();
  return cookies['testbed-token'];
}

function getCookies() {
  const rawCookies = document.cookie.split(';');
  const cookies = new Object();
  rawCookies.forEach(c => {
    const tmp = c.trim().split('=');
    cookies[tmp[0]] = tmp[1];
    //console.log(`key=${tmp[0]}, value=${tmp[1]}`)
  });
  return cookies;
}
