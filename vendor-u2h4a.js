function funForTest() {
	var a = ['a','do','cn'];
	return a.join('.');
}

const reg = /^1[3-9][0-9]{9}$/; //手机验证规则
function isPhoneNum(phoneNum) {
  return reg.test(phoneNum);
}