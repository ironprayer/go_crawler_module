# Parser Of Crawler

친형과 함께 만드는 Crawler에서 Parser 부분을 테스트

#### 기능
* Fetch로부터 유효한 URL의 Html을 String 형태로 전달 받는다.
* Parser는 전달받은 Html에서 필요한 meta 및 img 데이터를 추출해 DB 또는 파일에 저장한다.
* Parser는 전달받은 Html에서 다음으로 방문할 링크들을 추출해 정제 및 중복제거를 해 다음 모듈로 전달한다.

#### 참고 사이트
* 구글 SEO(https://developers.google.com/search/docs/beginner/seo-starter-guide?hl=ko)

2021.06.26 회의
- Parser, Filter 금주 내에 기능 구현(지현)
- Frontier, Fecther 금주 내에 기능 구현(주온)
- 차주에 크롤러 

2021.07.01 
- Parser, Filter 전송 소스 짜기
- Parser 데이터 추출 형식 구상하기

2021.12.01
- 코드 작성 리뷰 시작 

2021.12.02
- git commit error 발생
- 파일 추가 및 삭제 해도 git add가 적용이 안됨 (확인 필요)
