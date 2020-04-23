# Git

- Tạo tag

```bash
git tag v1.0.1
```

- Tạo Annotated tag

```bash
git tag -a v1.0.0 -m "tag description"
```

- Đẩy tag từ local lên remote

```bash
git push origin v1.0.0
```

- Xem danh sách tags

```bash
git tag
```

## Chú ý
Mỗi lần đẩy tag lên, GIT coi đó là 1 lần release