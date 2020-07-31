##### validator 
* go-playground/locales :多语言包，需要与universal-translator配套使用
* go-playground/universal-translator：通用翻译器，是一个使用CLDR数据+复数规则的语言i18n转换器
* go-playground/validator/v10/translations:validator的翻译器

在识别当前请求的语言类别时，我们通过GetHeader方法获取约定的header参数locale，判别当前请求的语言类别时en还是zh。
如果有其他语言环境要求，也可以继续引入其他语言类别，因为go-playground/locales支持所有语言类别

在后续的注册步骤中，调用RegisterDefaultTranslations方法，将验证器和对应语言类型的Translator注册进来，实现验证器的多语言支持。同时将Translator存储到全局上下文中，以便后续翻译时使用.
