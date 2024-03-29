# 인라인 준비

## 개요

데이터 준비의 전통적인 접근 방식은 원본 데이터 소스(일반적으로 로컬 파일 시스템의 폴더)를 32GiB보다 작은 여러 개의 CAR 파일로 변환하는 것입니다. 이렇게 하면 데이터 준비자에게 새로운 문제가 발생합니다. 1PiB의 데이터 세트를 준비하려면 CAR 파일을 저장할 추가적인 1PiB의 스토리지 서버를 제공해야 합니다.

인라인 준비는 이 문제를 해결하기 위해 CAR 파일의 블록을 원본 데이터 소스로 다시 매핑하여 내보낸 CAR 파일을 저장할 필요가 없도록 합니다.

## CAR 검색 작동 방식

인라인 준비를 사용하면 메타데이터 데이터베이스와 원본 데이터 소스를 사용하여 CAR 파일을 HTTP를 통해 서비스할 수 있습니다. 왜냐하면 인라인 준비가 특정 바이트 범위의 CAR 파일을 원본 데이터 소스로 다시 매핑할 수 있기 때문입니다.&#x20;

CAR 파일을 HTTP를 통해 서비스하려면 간단히 내용 공급자를 시작하고 선택적으로 리버스 프록시 뒤에 놓으면 됩니다.

```sh
singularity run content-provider
```

이렇게 하면 문제가 하나 더 발생합니다. 데이터 소스가 이미 원격 저장 시스템인 경우(예: S3 또는 FTP), 파일 내용은 여전히 싱귤러리 콘텐츠 공급자를 통해 스토리지 공급자로 프록시됩니다.

싱귤래리 메타데이터 API와 싱귤래리 다운로드 유틸리티를 사용하여 이를 해결할 수 있는 솔루션이 있습니다.

```bash
singularity run api
singularity download <piece_cid>
```

싱귤래리 메타데이터 API는 원본 데이터 소스와 싱귤래리 다운로드 유틸리티가 이러한 CAR 파일을 어떻게 조립할지에 대한 계획을 반환하며, 싱귤래리 다운로드 유틸리티는 이 계획을 해석하고 원본 데이터에서 로컬 CAR 파일로 스트리밍합니다. 중간에 변환이나 조립이 발생하지 않고 모든 작업이 스트림으로 이루어집니다.

메타데이터 API는 원본 데이터 소스에서 데이터에 액세스하기 위해 필요한 자격 증명을 반환하지 않으므로 스토리지 공급자는 자신의 데이터 소스에 액세스하고 해당 자격 증명을 싱귤래리 다운로드 명령에 제공해야 합니다.

## 오버헤드

인라인 준비는 매우 작은 오버헤드를 가지고 있습니다.

각 데이터 블록에 대해 데이터베이스 행을 저장해야 하므로 준비한 각 1MiB 블록에 대해 100바이트가 필요합니다. 1PiB 데이터 세트의 경우, 데이터베이스는 이러한 매핑 메타데이터를 저장하기 위해 10TiB의 디스크 공간이 필요합니다. 일반적으로 이는 문제가 되지 않지만, 작은 파일이 많은 데이터 세트의 경우 디스크 오버헤드가 크게 증가할 수 있습니다.

나중에 원본 데이터 소스에서 CAR 파일을 동적으로 다시 생성할 때는 데이터베이스에서 해당 매핑을 확인해야 합니다. 일반적으로 이는 문제가 되지 않습니다. 1GB/초의 대역폭은 데이터베이스에서 1000개의 항목을 조회하는 것으로 변환되며, 이는 모든 지원되는 데이터베이스 백엔드의 병목 지점 아래에 있으며, 향후 최적화가 이러한 오버헤드를 줄일 수 있습니다.

## 인라인 준비 활성화

인라인 준비는 암호화가 필요하지 않은 모든 데이터 세트에 대해 항상 활성화됩니다. 출력 디렉토리를 지정하여 데이터 세트를 생성하는 경우 CAR 파일이 해당 디렉토리에 내보내집니다. 이 경우 CAR 검색 요청은 먼저 해당 디렉토리에서 처리되며, 사용자가 이러한 CAR 파일을 삭제한 경우에만 원본 데이터 소스로 다시 전환됩니다.

## 관련 자원

[download.md](../cli-reference/download.md "mention")